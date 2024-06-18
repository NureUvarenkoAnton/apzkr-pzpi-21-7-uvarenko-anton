package com.uvarenko.petwalker.vm

import androidx.lifecycle.ViewModel
import com.uvarenko.petwalker.data.Pet
import com.uvarenko.petwalker.data.PetWalkerRepository
import com.uvarenko.petwalker.data.Walk
import com.uvarenko.petwalker.data.Walkers
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.SupervisorJob
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.collectLatest
import kotlinx.coroutines.launch
import kotlin.coroutines.CoroutineContext

class HomeScreenViewModel(
    private val repository: PetWalkerRepository
) : ViewModel(), CoroutineScope {

    override val coroutineContext: CoroutineContext
        get() = Dispatchers.Main + SupervisorJob()

    private val _pets = MutableStateFlow(emptyList<Pet>())
    val pets: StateFlow<List<Pet>> = _pets

    private val _walks = MutableStateFlow(emptyList<Walk>())
    val walks: StateFlow<List<Walk>> = _walks

    private val _walkers = MutableStateFlow(emptyList<Walkers>())
    val walkers: StateFlow<List<Walkers>> = _walkers

    private val _walkPosition = MutableStateFlow<Pair<Float, Float>?>(null)
    val walkPosition: StateFlow<Pair<Float, Float>?> = _walkPosition

    init {
        launch {
            repository.getPets()
            repository._pets.collectLatest {
                _pets.emit(it)
            }
        }
        launch {
            repository.getWalks()
            repository._walks.collectLatest {
                _walks.emit(it)
            }
        }
        launch {
            repository.getWalkers()
            repository._walkers.collectLatest {
                _walkers.emit(it)
            }
        }
        launch {
//            repository.getWalkers()
//            repository._walkers.collectLatest {
//                _walkers.emit(it)
//            }
        }
    }


    fun deletePet(pet: Pet) {
        launch {
            repository.deletePet(pet)
        }
    }

    fun add(pet: Pet) {
        launch {
            repository.addPet(pet)
        }
    }

    fun update(pet: Pet) {
        launch {
            repository.updatePet(pet)
        }
    }

    fun requestWalk(pet: Pet, walker: Walkers) {
        launch {
            repository.requestWalk(pet, walker)
        }
    }

}
