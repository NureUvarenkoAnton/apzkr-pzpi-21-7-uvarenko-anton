package com.uvarenko.petwalker.screens


import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.wrapContentHeight
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.rounded.Delete
import androidx.compose.material.icons.rounded.Edit
import androidx.compose.material3.Button
import androidx.compose.material3.Card
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
import androidx.compose.material3.TopAppBar
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import com.uvarenko.petwalker.data.Pet
import com.uvarenko.petwalker.vm.HomeScreenViewModel
import com.uvarenko.petwalker.vm.HomeScreenViewModelFactory

@Composable
fun PetsScreen(
    modifier: Modifier = Modifier,
    viewModel: HomeScreenViewModel,
    onRequestWalk: (Pet) -> Unit = {},
    onEdit: (Pet?) -> Unit = {},
) {
    val pets by viewModel.pets.collectAsState()
    Column(modifier = modifier) {
        Button(
            modifier = Modifier
                .fillMaxWidth()
                .padding(16.dp),
            onClick = { onEdit(null) }) {
            Text(text = "Add Pet")
        }
        LazyColumn(
            modifier = Modifier
                .fillMaxSize()
                .padding(16.dp)
        ) {
            items(pets.size) {
                PetItem(pets[it], onRequestWalk, onEdit, onDelete = {
                    viewModel.deletePet(it)
                })
                Spacer(modifier = Modifier.padding(top = 16.dp))
            }
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Preview
@Composable
fun PetItem(
    pet: Pet = Pet(1, "Vivi", 2, "happy little dog"),
    onRequestWalk: (Pet) -> Unit = {},
    onEdit: (Pet) -> Unit = {},
    onDelete: (Pet) -> Unit = {},
) {
    Card(
        shape = RoundedCornerShape(8.dp),
        onClick = { onRequestWalk(pet) }
    ) {
        Row(
            modifier = Modifier
                .fillMaxWidth()
                .wrapContentHeight()
                .padding(10.dp)
        ) {
            Column(
                modifier = Modifier
                    .wrapContentHeight()
                    .weight(1f)
            ) {
                Text(
                    text = pet.name ?: "",
                    style = MaterialTheme.typography.headlineLarge
                )
                Text(
                    text = "${pet.age} years old",
                    style = MaterialTheme.typography.titleMedium
                )
                Text(
                    text = pet.additionalInfo ?: "",
                    style = MaterialTheme.typography.bodyMedium
                )
            }
            Column(
                modifier = Modifier.fillMaxHeight(),
                verticalArrangement = Arrangement.SpaceEvenly
            ) {
                Icon(
                    modifier = Modifier
                        .size(30.dp)
                        .clip(CircleShape)
                        .clickable { onEdit(pet) },
                    imageVector = Icons.Rounded.Edit,
                    contentDescription = ""
                )
                Icon(
                    modifier = Modifier
                        .size(30.dp)
                        .clip(CircleShape)
                        .clickable { onDelete(pet) },
                    imageVector = Icons.Rounded.Delete,
                    contentDescription = ""
                )
            }
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun PetEditScreen(modifier: Modifier = Modifier, pet: Pet?, onBack: () -> Unit) {
    var name by remember { mutableStateOf(pet?.name ?: "") }
    var age by remember { mutableStateOf(pet?.age?.toString() ?: "") }
    var additionalInfo by remember { mutableStateOf(pet?.additionalInfo ?: "") }
    val context = LocalContext.current
    val viewModel =
        viewModel<HomeScreenViewModel>(factory = HomeScreenViewModelFactory(context))
    Scaffold(
        modifier = modifier,
        topBar = { TopAppBar(title = { Text("Add / Edit pet data") }) }
    ) {
        Column(
            modifier = Modifier
                .fillMaxSize()
                .padding(it)
                .padding(10.dp)
        ) {
            TextField(
                modifier = Modifier
                    .fillMaxWidth()
                    .padding(bottom = 20.dp),
                placeholder = { Text(text = "Name") },
                value = name,
                onValueChange = { name = it })
            TextField(
                modifier = Modifier
                    .fillMaxWidth()
                    .padding(bottom = 20.dp),
                keyboardOptions = KeyboardOptions.Default.copy(keyboardType = KeyboardType.Number),
                placeholder = { Text(text = "Age") },
                value = age,
                onValueChange = { age = it }
            )
            TextField(
                modifier = Modifier
                    .fillMaxWidth()
                    .padding(bottom = 20.dp),
                placeholder = { Text(text = "Additional info") },
                value = additionalInfo,
                onValueChange = { additionalInfo = it })
            Button(
                modifier = Modifier
                    .fillMaxWidth()
                    .padding(bottom = 20.dp),
                onClick = {
                    if (pet == null) {
                        viewModel.add(
                            Pet(
                                name = name,
                                age = age.toInt(),
                                additionalInfo = additionalInfo
                            )
                        )
                    } else {
                        viewModel.update(
                            pet.copy(
                                name = name,
                                age = age.toInt(),
                                additionalInfo = additionalInfo
                            )
                        )
                    }
                    onBack()
                }) {
                Text(text = "Save")
            }
        }
    }
}