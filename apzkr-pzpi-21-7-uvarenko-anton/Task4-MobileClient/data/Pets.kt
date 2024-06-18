package com.uvarenko.petwalker.data

import android.os.Parcelable
import kotlinx.parcelize.Parcelize
import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable
import kotlinx.serialization.json.JsonNames

@Parcelize
@Serializable
data class Pet(
    @JsonNames("petId")
    @SerialName("pet_id")
    val id: Int? = null,
    @JsonNames("petName")
    @SerialName("name")
    val name: String? = null,
    @JsonNames("age")
    @SerialName("age")
    val age: Int,
    @JsonNames("PetAdditionalInfo", "additionalInfo")
    @SerialName("additionalInfo")
    val additionalInfo: String? = null,
) : Parcelable
