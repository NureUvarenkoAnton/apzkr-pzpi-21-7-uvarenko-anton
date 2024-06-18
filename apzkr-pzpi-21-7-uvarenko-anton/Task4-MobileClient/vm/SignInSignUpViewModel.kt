package com.uvarenko.petwalker.vm

import androidx.lifecycle.ViewModel
import com.uvarenko.petwalker.data.PetWalkerRepository
import com.uvarenko.petwalker.data.SignInSignUp
import com.uvarenko.petwalker.data.UserType
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.SupervisorJob
import kotlinx.coroutines.launch
import kotlin.coroutines.CoroutineContext

class SignInSignUpViewModel(
    private val repository: PetWalkerRepository
) : ViewModel(), CoroutineScope {

    override val coroutineContext: CoroutineContext
        get() = Dispatchers.Main + SupervisorJob()

    fun signIn(email: String, password: String) {
        launch {
            repository.signIn(
                SignInSignUp(
                    email = email,
                    password = password
                )
            )
        }
    }

    fun signUp(email: String, password: String, name: String, type: UserType) {
        launch {
            repository.signUp(
                SignInSignUp(
                    email = email,
                    password = password,
                    name = name,
                    userType = type
                )
            )
        }
    }

}
