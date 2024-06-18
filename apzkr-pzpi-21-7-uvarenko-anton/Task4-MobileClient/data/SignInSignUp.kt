package com.uvarenko.petwalker.data

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class SignInSignUp(
    @SerialName("email")
    val email: String,
    @SerialName("password")
    val password: String,
    @SerialName("name")
    val name: String? = null,
    @SerialName("userType")
    val userType: UserType? = null
)

enum class UserType(val value: String) {
    WALKER("walker"),
    DEFAULT("default")
}


@Serializable
data class SignInSignUpResponse(
    @SerialName("token")
    val token: String
)