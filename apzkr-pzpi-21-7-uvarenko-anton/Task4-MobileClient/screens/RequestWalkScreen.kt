package com.uvarenko.petwalker.screens

import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.wrapContentHeight
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Card
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import com.uvarenko.petwalker.data.Pet
import com.uvarenko.petwalker.data.Walkers
import com.uvarenko.petwalker.vm.HomeScreenViewModel
import com.uvarenko.petwalker.vm.HomeScreenViewModelFactory

@Composable
fun RequestWalkScreen(modifier: Modifier = Modifier, pet: Pet) {
    val context = LocalContext.current
    val viewModel = viewModel<HomeScreenViewModel>(factory = HomeScreenViewModelFactory(context))
    Scaffold(
        modifier = modifier,
    ) {
        val walkers by viewModel.walkers.collectAsState(initial = emptyList())
        LazyColumn(
            modifier = Modifier.padding(it)
        ) {
            items(walkers.size) {
                Walker(
                    modifier = Modifier
                        .fillMaxWidth()
                        .wrapContentHeight()
                        .padding(horizontal = 10.dp)
                        .clickable { viewModel.requestWalk(pet, walkers[it]) },
                    walker = walkers[it]
                )
                Spacer(modifier = Modifier.height(10.dp))
            }
        }
    }
}

@Composable
fun Walker(modifier: Modifier = Modifier, walker: Walkers) {
    Card(
        modifier = modifier,
        shape = RoundedCornerShape(8.dp)
    ) {
        Column(
            modifier = Modifier
                .fillMaxWidth()
                .padding(10.dp)
        ) {
            Text(
                text = walker.name,
                style = MaterialTheme.typography.headlineLarge
            )
            Text(
                text = walker.email,
                style = MaterialTheme.typography.headlineMedium
            )
        }
    }
}