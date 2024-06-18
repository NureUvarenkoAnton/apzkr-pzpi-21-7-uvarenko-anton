package com.uvarenko.petwalker.vm

import android.content.Context
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import com.uvarenko.petwalker.data.PetWalkerRepository

class SignInSignUpViewModelFactory(private val context: Context) : ViewModelProvider.Factory {

    @Suppress("UNCHECKED_CAST")
    override fun <T : ViewModel> create(modelClass: Class<T>): T =
        SignInSignUpViewModel(PetWalkerRepository.getInstance(context)) as T

}
