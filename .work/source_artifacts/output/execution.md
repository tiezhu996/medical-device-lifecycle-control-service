# ld-336 医疗器械资产管理平台 — 执行验证报告

- 项目编号：ld-336（短名 medasset，全栈：Angular 17 + Go 1.22/Gin/GORM + MySQL 8 + Redis 7）
- 项目路径：`/Users/gaobo/repositories/gitlab/评审项目/0-1代码生成提示词/golang-改编提示词/医疗健康主题项目提示词/ld-336`
- 验证时间：2026-08-17（Asia/Shanghai）
- 端口：前端 18936 → 80，后端 19936 → 8080，DB 44017 → 3306，Redis 46317 → 6379（MinIO 无）

## 启动命令

```bash
docker compose config --quiet        # 通过
docker compose up -d --build         # 一键启动
docker compose ps                    # 全部 healthy
```

## docker compose ps（healthy 结果）

```
NAME                IMAGE               COMMAND                  SERVICE    CREATED          STATUS                    PORTS
medasset-backend    medasset-backend    "/app/server"            backend    7 minutes ago    Up 7 minutes (healthy)    0.0.0.0:19936->8080/tcp, [::]:19936->8080/tcp
medasset-db         mysql:8.0           "docker-entrypoint.s…"   db         11 minutes ago   Up 11 minutes (healthy)   33060/tcp, 0.0.0.0:44017->3306/tcp, [::]:44017->3306/tcp
medasset-frontend   medasset-frontend   "/docker-entrypoint.…"   frontend   7 minutes ago    Up 7 minutes              0.0.0.0:18936->80/tcp, [::]:18936->80/tcp
medasset-redis      redis:7-alpine      "docker-entrypoint.s…"   redis      11 minutes ago   Up 11 minutes (healthy)   0.0.0.0:46317->6379/tcp, [::]:46317->6379/tcp
```

## curl 接口冒烟清单（≥8 项，实际 20 项，全部通过）

| # | 方法 | 路径 | 状态码/结果 |
| --- | --- | --- | --- |
| 1 | GET | /healthz | 200 `{"code":0,...}` |
| 2 | POST | /api/v1/auth/login | 200，返回 JWT（SUPER_ADMIN） |
| 3 | GET | /api/v1/auth/me | 200，返回当前用户 |
| 4 | POST | /api/v1/devices | 200，生成设备 MA-0001 及分发条码 BAR-MA-0001-* |
| 5 | GET | /api/v1/devices?page=1&page_size=10 | 200，total=1 |
| 6 | POST | /api/v1/purchases | 200，pending_device_admin |
| 7 | POST | /api/v1/purchases/1/device-admin-approve | 200，pending_dean |
| 8 | POST | /api/v1/purchases/1/dean-approve | 200，approved |
| 9 | POST | /api/v1/purchases/1/deliver | 200，delivered |
| 10 | POST | /api/v1/purchases/1/accept | 200，生成设备 MA-0002 及条码 |
| 11 | POST | /api/v1/maintenances | 200，报修工单 pending |
| 12 | POST | /api/v1/maintenances/1/start | 200，in_progress（设备→维修中） |
| 13 | POST | /api/v1/maintenances/1/complete | 200，completed（cost=1500） |
| 14 | POST | /api/v1/calibrations | 200，计量台账 normal |
| 15 | POST | /api/v1/calibrations/1/result（不合格） | 200，unqualified 且设备自动禁用 |
| 16 | POST | /api/v1/transfers + /transfers/1/approve | 200，approved（设备科室自动变更为急诊科） |
| 17 | POST | /api/v1/scraps + /scraps/1/approve | 200，approved（设备状态→scrapped） |
| 18 | GET | /api/v1/stats/overview | 200，total_devices=2, total_amount=3280000, maintenance_cost=1500 |
| 19 | GET | /api/v1/audits?page=1&page_size=5 | 200，audit_total=17 |
| 20 | POST | /api/v1/maintenances/plan/generate | 200，自动生成 4 条保养计划 |

## 浏览器验证（内置 playwright 包装脚本，独立 --session，无外部 Chrome）

页面与截图（`<项目目录>/output/`）：

| 页面 | 验证内容 | 截图 |
| --- | --- | --- |
| 登录 | admin/admin123 登录成功 | output/01-login.png |
| 资产总览 | 设备总数/资产总值等统计卡片渲染 | output/02-dashboard.png |
| 设备台账 | 列表展示 4 台设备；「登记设备」弹窗创建 MA-0003 监护仪成功 | output/03-devices.png |
| 采购验收 | 列表展示采购单；「提交申请」创建呼吸机；UI 走完 设备科通过→院长通过→到货登记→验收 全流程（最终已验收） | output/04-purchases.png |
| 维护保养 | 自动生成保养计划；报修/创建工单入口 | output/05-maintenance.png |
| 计量质控 | 计量台账 JL-0001 不合格展示 | output/06-calibrations.png |
| 设备调拨 | 调拨单 TR2026* 已批准 | output/07-transfers.png |
| 报废管理 | 报废单 SC2026* 已批准 | output/08-scraps.png |
| 审计日志 | 操作人/动作/模块/详情等审计记录展示 | output/09-audits.png |

浏览器交互覆盖：登录、侧边栏路由、设备创建弹窗、采购申请弹窗、审批确认弹窗、验收弹窗、保养计划生成按钮、退出登录。全部页面无阻断性控制台错误。

## 修复过程摘要

1. **事务内更新使用 base DB 导致锁等待超时**：service 事务闭包内调用 `repo.Update`（走 `r.db`），与 `SELECT ... FOR UPDATE` 锁冲突（MySQL Error 1205）。为各仓储新增 `UpdateTx(tx, entity)` 并替换事务内调用。
2. **MySQL 保留字 `key` 导致统计 SQL 报错**：`GroupCount` 的 `AS key` 改为 `AS group_key`。
3. **`/api/v1/auth/me` 未挂载认证中间件**：拆分为公开（register/login）与受保护（me）两组路由。
4. **前端 filter-bar 的 `formControlName` 未包裹 `[formGroup]`**：Angular 报 `Cannot read properties of null (reading 'addControl')`，为 4 个页面（设备/采购/保养/审计）补 `<form [formGroup]>`。
5. 补充 `MatCardModule` / `MatIconModule` / `inject` 导入缺失导致的编译错误。

## git commit hash

- 主实现提交：`0a84861ac9f5c3703123542b2db2922fac069ddb`
- 最终提交：`（见下方 git log，报告回填提交）`
