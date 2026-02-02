package com.grfsv.palbum.core.datastore

import androidx.datastore.core.Serializer
import com.google.crypto.tink.Aead
import java.io.InputStream
import java.io.OutputStream
import java.security.GeneralSecurityException
import javax.inject.Inject

class RefreshTokenSerializer @Inject constructor(
   private val aead: Aead
) : Serializer<ByteArray> {
    override val defaultValue: ByteArray = ByteArray(0)

    override suspend fun readFrom(input: InputStream): ByteArray {
        val bytes = input.readBytes()

        if (bytes.isEmpty()) return defaultValue

        return try {
            aead.decrypt(bytes, ByteArray(0))
        } catch (e : GeneralSecurityException) {
            defaultValue
        }
    }

    override suspend fun writeTo(t: ByteArray, output: OutputStream) {
        if (t.isEmpty()) {
            output.write(t)
            return
        }
        output.write(aead.encrypt(t, ByteArray(0)))
    }

}