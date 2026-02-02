# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## プロジェクト概要

Palbumは、モダンなAndroid開発のベストプラクティスに基づいた、モジュール化されたマルチレイヤーアーキテクチャを採用したAndroidアプリケーションです。Jetpack Compose、MVVM、Kotlin Serializationを使用した100% Kotlin実装です。

## ビルドとテストコマンド

### ビルド
```bash
# プロジェクト全体をビルド
./gradlew build

# アプリモジュールのみをビルド
./gradlew :app:build

# リリースビルド
./gradlew assembleRelease

# デバッグビルド
./gradlew assembleDebug

# プロジェクトをクリーン
./gradlew clean
```

### テスト
```bash
# すべてのユニットテストを実行
./gradlew test

# アプリモジュールのユニットテストのみ実行
./gradlew :app:test

# すべてのInstrumentedテストを実行（エミュレータ/実機が必要）
./gradlew connectedAndroidTest

# 特定のモジュールのInstrumentedテストを実行
./gradlew :app:connectedAndroidTest

# 特定のテストクラスを実行
./gradlew :app:test --tests "com.grfsv.palbum.ExampleUnitTest"
```

### コード品質
```bash
# KtLintでコードをチェック
./gradlew ktlintCheck

# KtLintで自動フォーマット
./gradlew ktlintFormat
```

### その他
```bash
# 依存関係ツリーを確認
./gradlew :app:dependencies

# タスク一覧を表示
./gradlew tasks
```

## プロジェクト構造

### モジュール構成
```
Palbum/
├── app/                    # メインアプリケーションモジュール（単一Activity）
├── core/
│   └── design/            # デザインシステム（テーマ、色、タイポグラフィ）
├── feature/
│   ├── home/              # ホーム機能モジュール
│   └── setting/           # 設定機能モジュール
└── build-logic/
    └── convention/        # Gradle Convention Plugins（共通ビルド設定）
```

### 技術スタック
- **Min SDK**: 36 (Android 15)
- **Target SDK**: 36
- **Kotlin**: 2.2.20
- **Java**: 21
- **Compose BOM**: 2025.10.00
- **Navigation Compose**: 2.9.5
- **Kotlin Serialization**: 1.9.0
- **Hilt**: 2.57.1（定義済みだが現在未使用）

## アーキテクチャパターン

### MVVM（Model-View-ViewModel）
- ViewModelは `androidx.lifecycle.ViewModel` を継承
- `StateFlow<UiState>` で UI状態を管理
- sealed class で型安全な状態定義（Idle, Loading, Success, Error）
- Compose UI は `collectAsState()` で StateFlow を監視

例: `feature/home/src/main/kotlin/com/grfsv/palbum/feature/home/HomeViewModel.kt`

### 単一Activity設計
- `app/src/main/kotlin/com/grfsv/palbum/MainActivity.kt` のみ
- すべての画面はJetpack Composeで実装
- Edge-to-Edge UI対応

### ナビゲーション
- Jetpack Navigation Compose使用
- Kotlin Serializationで型安全なルート定義（`@Serializable`）
- ルート定義: `app/src/main/kotlin/com/grfsv/palbum/navigation/`
- Bottom Navigation Barあり

例:
```kotlin
@Serializable
object HomeRoute

@Serializable
data class OrderDetailRoute(val orderId: String)
```

## モジュール依存関係

```
app
├── :core:design
├── :feature:home
└── :feature:setting

feature:home
├── AndroidX Compose
├── AndroidX Navigation
└── Material3

core:design
├── AndroidX Compose
├── Material3
└── Google Fonts (Noto Sans JP)
```

## Gradle Convention Plugins

`build-logic/convention/` に共通ビルド設定を定義:

- `com.grfsv.palbum.buildLogic.android.application` - アプリケーション用
- `com.grfsv.palbum.buildLogic.android.library` - ライブラリ用

共通設定:
- Android SDK設定（compileSdk, minSdk, targetSdk）
- Kotlin コンパイラ設定（Java 21 toolchain）
- KtLint統合
- BuildConfig生成

## Version Catalog

すべての依存関係は `gradle/libs.versions.toml` で一元管理。

モジュールで使用する場合:
```kotlin
implementation(libs.androidx.core.ktx)
implementation(libs.androidx.compose.bom)
```

## デザインシステム

`core:design` モジュールで管理:

- `AppTheme()` - Material3テーマ
- Dynamic Color対応（Android 12+）
- Light/Dark/Medium Contrast/High Contrastモード
- Google Fonts統合（Noto Sans JP）

すべてのCompose画面は `AppTheme { }` でラップ。

## 今後の実装予定

- Hilt DIの統合（定義済み、未使用）
- データレイヤー（Repository/DataSource）の実装
- ネットワーク層（Retrofit/OkHttpなど）
- 本格的なテストカバレッジ
- `feature:setting` の完成

## namespace の分離

各モジュールは独自のnamespaceを持つ:

- `app`: `com.grfsv.palbum`
- `core:design`: `com.grfsv.palbum.core.design`
- `feature:home`: `com.grfsv.palbum.feature.home`
- `feature:setting`: `com.grfsv.palbum.feature.setting`
- `build-logic`: `com.grfsv.palbum.buildLogic`

## 注意事項

- 新しい機能を追加する場合、`feature/` 配下に新しいモジュールを作成
- UIコンポーネントは100% Jetpack Compose（XML不使用）
- 状態管理は必ずViewModelのStateFlowで行う
- ナビゲーションルートは `@Serializable` で型安全に定義
- コード品質のためKtLintを必ず実行
- 新しいモジュールは `settings.gradle.kts` に追加が必要