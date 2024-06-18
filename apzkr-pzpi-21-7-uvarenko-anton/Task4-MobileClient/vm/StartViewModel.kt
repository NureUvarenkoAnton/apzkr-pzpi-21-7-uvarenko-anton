package com.uvarenko.petwalker.vm

import androidx.lifecycle.ViewModel
import com.uvarenko.petwalker.data.LocalStore
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.SupervisorJob
import kotlinx.coroutines.delay
import kotlinx.coroutines.flow.MutableSharedFlow
import kotlinx.coroutines.flow.SharedFlow
import kotlinx.coroutines.launch
import kotlin.coroutines.CoroutineContext

class StartViewModel(localStore: LocalStore) : ViewModel(), CoroutineScope {

    override val coroutineContext: CoroutineContext
        get() = Dispatchers.Main + SupervisorJob()

    private val _initNav = MutableSharedFlow<String>()
    val initNav: SharedFlow<String> = _initNav

    init {
        launch {
            delay(1_000)
            if (localStore.getToken() == null) {
                _initNav.emit("signin")
            } else {
                _initNav.emit("home")
            }
        }
    }

}
