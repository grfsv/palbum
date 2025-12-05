# Remind Map API 仕様書

## 概要

Remind Map APIは、ユーザー認証とトークン管理機能を提供するRESTful APIです。このAPIは、ユーザーのサインアップ、ログイン、ログアウト、およびトークンのリフレッシュ機能を提供します。

### ベースURL

```text
http://localhost:8080
```

本番環境では、適切なドメインとHTTPSプロトコルを使用してください。

### 共通ヘッダー

すべてのリクエストには以下のヘッダーを含めてください：

```text
Content-Type: application/json
```

### 認証

認証が必要なエンドポイントでは、Bearer Token認証を使用します。

```text
Authorization: Bearer <accessToken>
```

アクセストークンは、サインアップまたはログイン時に取得できます。トークンの有効期限が切れた場合は、リフレッシュトークンを使用して新しいアクセストークンを取得してください。

---

## エンドポイント一覧

### 1. ヘルスチェック

APIサーバーの稼働状態を確認するためのエンドポイントです。

#### リクエスト

```text
GET /health
```

#### 認証

不要

#### リクエストパラメータ

なし

#### レスポンス

**ステータスコード**: `200 OK`

**レスポンスボディ**:

```json
{
  "status": "ok",
  "message": "Remind Map API is running started"
}
```

#### フィールド説明

| フィールド | 型 | 説明 |
|----------|-----|------|
| status | string | サーバーの状態（常に "ok"） |
| message | string | ステータスメッセージ |

---

### 2. ユーザーサインアップ

新規ユーザーを登録し、アクセストークンとリフレッシュトークンを取得します。

#### リクエスト

```text
POST /v1/auth/signup
```

#### 認証

不要

#### リクエストボディ

```json
{
  "name": "田中太郎",
  "mail": "tanaka@example.com",
  "password": "securePassword123"
}
```

#### リクエストフィールド

| フィールド | 型 | 必須 | 説明 |
|----------|-----|------|------|
| name | string | Yes | ユーザー名 |
| mail | string | Yes | メールアドレス |
| password | string | Yes | パスワード |

#### レスポンス

**ステータスコード**: `201 Created`

**レスポンスボディ**:

```json
{
  "data": {
    "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "name": "田中太郎",
    "mail": "tanaka@example.com"
  }
}
```

#### レスポンスフィールド

| フィールド | 型 | 説明 |
|----------|-----|------|
| data.accessToken | string | APIアクセス用のJWTトークン |
| data.refreshToken | string | アクセストークン更新用のJWTトークン |
| data.name | string | 登録されたユーザー名 |
| data.mail | string | 登録されたメールアドレス |

#### エラーレスポンス

**409 Conflict** - メールアドレスが既に使用されている場合

```json
{
  "error": "email already in use"
}
```

**400 Bad Request** - バリデーションエラー

```json
{
  "error": "バリデーションエラーメッセージ"
}
```

または、複数のバリデーションエラーがある場合：

```json
{
  "errors": [
    "エラーメッセージ1",
    "エラーメッセージ2"
  ]
}
```

---

### 3. ユーザーログイン

既存のユーザーでログインし、アクセストークンとリフレッシュトークンを取得します。

#### リクエスト

```text
POST /v1/auth/login
```

#### 認証

不要

#### リクエストボディ

```json
{
  "mail": "tanaka@example.com",
  "password": "securePassword123"
}
```

#### リクエストフィールド

| フィールド | 型 | 必須 | 説明 |
|----------|-----|------|------|
| mail | string | Yes | 登録済みのメールアドレス |
| password | string | Yes | パスワード |

#### レスポンス

**ステータスコード**: `200 OK`

**レスポンスボディ**:

```json
{
  "data": {
    "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

#### レスポンスフィールド

| フィールド | 型 | 説明 |
|----------|-----|------|
| data.accessToken | string | APIアクセス用のJWTトークン |
| data.refreshToken | string | アクセストークン更新用のJWTトークン |

#### エラーレスポンス

**401 Unauthorized** - メールアドレスが見つからない、またはパスワードが不正な場合

```json
{
  "error": "unauthorized"
}
```

**400 Bad Request** - バリデーションエラー

```json
{
  "error": "バリデーションエラーメッセージ"
}
```

---

### 4. ユーザーログアウト

リフレッシュトークンを無効化してログアウトします。

#### リクエスト

```text
POST /v1/auth/logout
```

#### 認証

不要（ただし、リフレッシュトークンが必要）

#### リクエストボディ

```json
{
  "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### リクエストフィールド

| フィールド | 型 | 必須 | 説明 |
|----------|-----|------|------|
| refreshToken | string | Yes | ログアウトするリフレッシュトークン |

#### レスポンス

**ステータスコード**: `204 No Content`

レスポンスボディはありません。

#### エラーレスポンス

**401 Unauthorized** - トークンが無効または見つからない場合

```json
{
  "error": "unauthorized"
}
```

**400 Bad Request** - バリデーションエラー（必須フィールドが欠けている場合）

```json
{
  "error": "バリデーションエラーメッセージ"
}
```

---

### 5. トークンリフレッシュ

リフレッシュトークンを使用して、新しいアクセストークンとリフレッシュトークンを取得します。

#### リクエスト

```text
POST /v1/auth/refresh
```

#### 認証

不要（ただし、有効なリフレッシュトークンが必要）

#### リクエストボディ

```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### リクエストフィールド

| フィールド | 型 | 必須 | 説明 |
|----------|-----|------|------|
| refresh_token | string | Yes | 有効なリフレッシュトークン |

**注意**: このエンドポイントでは、フィールド名が `refresh_token`（スネークケース）です。他のエンドポイントでは `refreshToken`（キャメルケース）を使用しています。

#### レスポンス

**ステータスコード**: `200 OK`

**レスポンスボディ**:

```json
{
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

#### レスポンスフィールド

| フィールド | 型 | 説明 |
|----------|-----|------|
| data.accessToken | string | 新しいAPIアクセス用のJWTトークン |
| data.refreshToken | string | 新しいリフレッシュトークン |

#### エラーレスポンス

**401 Unauthorized** - トークンが無効な場合

```json
{
  "error": "unauthorized"
}
```

**400 Bad Request** - バリデーションエラー（必須フィールドが欠けている場合）

```json
{
  "error": "バリデーションエラーメッセージ"
}
```

---

## エラーレスポンス形式

このAPIでは、エラーが発生した場合に以下の形式でレスポンスを返します。

### 単一エラー

バリデーションエラーや認証エラーなど、単一のエラーメッセージを返す場合：

```json
{
  "error": "エラーメッセージ"
}
```

### 複数バリデーションエラー

複数のバリデーションエラーがある場合、エラーメッセージの配列を返します：

```json
{
  "errors": [
    "エラーメッセージ1",
    "エラーメッセージ2",
    "エラーメッセージ3"
  ]
}
```

---

## バージョン情報

- **APIバージョン**: v1
- **最終更新日**: 2025-12-05
