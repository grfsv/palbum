plugins {
    alias(libs.plugins.convention.hilt)
    alias(libs.plugins.kotlin.serialization)
    alias(libs.plugins.convention.android.library)
}

android {
    namespace = "com.grfsv.palbum.core.model"


}

dependencies {
    implementation(libs.tink)
    implementation(libs.kotlinx.serialization.json)
}