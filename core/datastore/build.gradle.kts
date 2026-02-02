plugins {
    alias(libs.plugins.convention.hilt)
    alias(libs.plugins.kotlin.serialization)
    alias(libs.plugins.convention.android.library)
}

android {
    namespace = "com.grfsv.palbum.core.datastore"

}

dependencies {
    implementation(project(":core:model"))
    implementation(libs.tink)
    implementation("androidx.datastore:datastore-core:1.0.0")

}