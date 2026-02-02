plugins {
    alias(libs.plugins.convention.android.library)
    alias(libs.plugins.convention.room)
    alias(libs.plugins.convention.hilt)
}

android {
    namespace = "com.grfsv.palbum.core.database"
}

dependencies {
    implementation(libs.kotlinx.datetime)
}