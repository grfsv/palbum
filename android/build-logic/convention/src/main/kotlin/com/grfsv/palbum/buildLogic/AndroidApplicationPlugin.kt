import com.android.build.api.dsl.ApplicationExtension
import com.grfsv.palbum.buildLogic.convention.android
import com.grfsv.palbum.buildLogic.convention.configureAndroidCommon
import com.grfsv.palbum.buildLogic.convention.configureKotlinCommon
import com.grfsv.palbum.buildLogic.convention.configureLinter
import com.grfsv.palbum.buildLogic.convention.getPluginId
import com.grfsv.palbum.buildLogic.convention.kotlinAndroid
import com.grfsv.palbum.buildLogic.convention.libs
import org.gradle.api.Plugin
import org.gradle.api.Project


class AndroidApplicationPlugin : Plugin<Project> {
    override fun apply(target: Project) {
        with(target) {
            with(pluginManager) {
                apply(libs.getPluginId("android-application"))
                apply(libs.getPluginId("kotlin-android"))
            }

            kotlinAndroid {
                configureKotlinCommon(this)
            }

            android<ApplicationExtension> {
                configureAndroidCommon()
                defaultConfig {
                    targetSdk = 36
                }
            }
            configureLinter()
        }
    }
}