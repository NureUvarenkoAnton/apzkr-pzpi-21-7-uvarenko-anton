package com.uvarenko.petwalker.data

import android.os.Parcelable
import kotlinx.parcelize.Parcelize
import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Parcelize
@Serializable
data class Walk(
    @SerialName("walkId")
    val id: Int,
    @SerialName("walkState")
    val walkState: String?,
    @SerialName("startTime")
    val startTime: String?,
    @SerialName("finishTime")
    val finishTime: String? = null,
    @SerialName("ownerId")
    val ownerId: Int?,
    @SerialName("ownerName")
    val ownerName: String?,
    @SerialName("ownerEmail")
    val ownerEmail: String?,
    @SerialName("walkerId")
    val walkerId: Int,
    @SerialName("walkerName")
    val walkerName: String,
    @SerialName("walkerEmail")
    val walkerEmail: String?,
    @SerialName("petId")
    val petId: Int,
    @SerialName("petName")
    val petName: String,
    @SerialName("PetAdditionalInfo")
    val petAdditionalInfo: String? = null
) : Parcelable