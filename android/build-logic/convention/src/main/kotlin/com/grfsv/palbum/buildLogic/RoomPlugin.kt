import com.google.devtools.ksp.gradle.KspExtension
import com.grfsv.palbum.buildLogic.convention.getPluginId
import com.grfsv.palbum.buildLogic.convention.libs
import org.gradle.api.Plugin
import org.gradle.api.Project
import org.gradle.kotlin.dsl.apply
import org.gradle.kotlin.dsl.configure
import org.gradle.kotlin.dsl.dependencies

class RoomPlugin : Plugin<Project> {
    override fun apply(target: Project) {
        with(target) {
            apply(plugin = libs.getPluginId("ksp"))

            pluginManager.withPlugin("com.android.base") {
                extensions.configure<KspExtension> {
                    arg("room.schemaLocation", "$projectDir/schemas")
                }

                dependencies {
                    "implementation"(libs.findLibrary("room-runtime").get())
                    "ksp"(libs.findLibrary("room-compiler").get())
                }
            }
        }
    }
}