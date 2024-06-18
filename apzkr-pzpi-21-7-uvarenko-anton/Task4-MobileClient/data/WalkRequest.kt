package com.uvarenko.petwalker.data

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class WalkRequest(
    @SerialName("walkerId")
    val walkerId: Int,
    @SerialName("petId")
    val petId: Int,
    @SerialName("startTime")
    val startTime: String,
)