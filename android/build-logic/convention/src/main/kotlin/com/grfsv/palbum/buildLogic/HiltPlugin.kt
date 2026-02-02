import com.grfsv.palbum.buildLogic.convention.getPluginId
import com.grfsv.palbum.buildLogic.convention.libs
import org.gradle.api.Plugin
import org.gradle.api.Project
import org.gradle.kotlin.dsl.apply
import org.gradle.kotlin.dsl.dependencies

class HiltPlugin : Plugin<Project> {
    override fun apply(target: Project) {
        with(target) {
            apply(plugin = libs.getPluginId("ksp"))

            pluginManager.withPlugin("com.android.base") {
                apply(plugin = libs.getPluginId("hilt"))

                // 2. 依存関係を追加
                dependencies {
                    // コンパイラ（ksp）の追加
                    "ksp"(libs.findLibrary("hilt-compiler").get())

                    // 実行時ライブラリ（implementation）の追加
                    "implementation"(libs.findLibrary("hilt-android").get())
                }
            }
        }
    }
}