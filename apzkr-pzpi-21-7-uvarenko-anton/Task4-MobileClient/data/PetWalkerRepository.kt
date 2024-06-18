package com.uvarenko.petwalker.data

import android.content.Context
import android.util.Log
import com.uvarenko.petwalker.network.baseUrl
import com.uvarenko.petwalker.network.ktorHttpClient
import io.ktor.client.HttpClient
import io.ktor.client.features.get
import io.ktor.client.request.delete
import io.ktor.client.request.get
import io.ktor.client.request.headers
import io.ktor.client.request.parameter
import io.ktor.client.request.post
import io.ktor.client.request.put
import kotlinx.coroutines.delay
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.serialization.decodeFromString
import kotlinx.serialization.json.Json
import java.text.SimpleDateFormat
import java.time.format.DateTimeFormatter
import java.util.Date

class PetWalkerRepository private constructor(
    private val networkClient: HttpClient,
    private val localStore: LocalStore
) {

    private val json = Json { ignoreUnknownKeys = true }

    val _pets = MutableStateFlow(emptyList<Pet>())
    val _walks = MutableStateFlow(emptyList<Walk>())
    val _walkers = MutableStateFlow(emptyList<Walkers>())

    companion object {

        private var INSTANCE: PetWalkerRepository? = null

        fun getInstance(context: Context): PetWalkerRepository {
            if (INSTANCE == null) {
                INSTANCE = PetWalkerRepository(ktorHttpClient, LocalStore(context))
            }
            return INSTANCE
                ?: throw IllegalStateException("could not initialize PetWalkerRepository")
        }
    }

    suspend fun signUp(data: SignInSignUp) {
        val response = networkClient.post<SignInSignUpResponse>("${baseUrl}/auth/register") {
            body = data
        }
        localStore.saveToken(response.token)
    }

    suspend fun signIn(data: SignInSignUp) {
        val response = networkClient.post<SignInSignUpResponse>("${baseUrl}/auth/login") {
            body = data
        }
        localStore.saveToken(response.token)
    }

    suspend fun addPet(data: Pet) {
        networkClient.post<String>("${baseUrl}/profile/pet") {
            body = data
            headers {
                localStore.getToken()?.let { append("Authorization", "Bearer $it") }
            }
        }
        getPets()
    }

    suspend fun updatePet(data: Pet) {
        networkClient.put<String>("${baseUrl}/profile/pet") {
            body = data
            headers {
                localStore.getToken()?.let { append("Authorization", "Bearer $it") }
            }
        }
        getPets()
    }

    suspend fun deletePet(data: Pet) {
        networkClient.delete<String>("${baseUrl}/profile/pet/${data.id}") {
            headers {
                localStore.getToken()?.let { append("Authorization", "Bearer $it") }
            }
        }
        getPets()
    }

    suspend fun requestWalk(pet: Pet, walkers: Walkers) {
        networkClient.post<String?>("${baseUrl}/walk/") {
            body = WalkRequest(
                walkerId = walkers.id,
                petId = pet.id!!,
                startTime = SimpleDateFormat("yyyy-MM-dd hh:mm:ss").format(Date())
            )
            headers {
                localStore.getToken()?.let { append("Authorization", "Bearer $it") }
            }
        }
        delay(500)
        getWalks()
    }

    suspend fun getPets() {
        val result = networkClient.get<String?>("${baseUrl}/profile/pets/EN") {
            headers {
                localStore.getToken()?.let { append("Authorization", "Bearer $it") }
            }
        }?.let { kotlin.runCatching { json.decodeFromString<List<Pet>>(it) }.getOrNull() }
            ?: emptyList()
        _pets.emit(result)
    }

    suspend fun getWalks() {
        val result = networkClient.get<String?>("${baseUrl}/walk/self/EN") {
            headers {
                localStore.getToken()?.let { append("Authorization", "Bearer $it") }
            }
        }?.let { kotlin.runCatching { json.decodeFromString<List<Walk>>(it) }.getOrNull() }
            ?: emptyList()
        _walks.emit(result)
    }

    suspend fun getWalkers() {
        val result = networkClient.get<String?>("${baseUrl}/users/walkers") {
            headers {
                localStore.getToken()?.let { append("Authorization", "Bearer $it") }
            }
        }?.let { kotlin.runCatching { json.decodeFromString<List<Walkers>>(it) }.getOrNull() }
            ?: emptyList()
        _walkers.emit(result)
    }

    suspend fun getWalkPosition(walk: Walk) {

    }

}