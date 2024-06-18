package com.uvarenko.petwalker.data

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class Walkers(
    @SerialName("id")
    val id: Int,
    @SerialName("name")
    val name: String,
    @SerialName("email")
    val email: String
)
