package com.uvarenko.petwalker.screens

import androidx.compose.foundation.clickable
import androidx.compose.foundation.interaction.MutableInteractionSource
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Home
import androidx.compose.material.icons.filled.List
import androidx.compose.material.icons.filled.Share
import androidx.compose.material.icons.outlined.Home
import androidx.compose.material.icons.outlined.List
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.uvarenko.petwalker.data.Pet
import com.uvarenko.petwalker.data.Walk
import com.uvarenko.petwalker.vm.HomeScreenViewModel
import com.uvarenko.petwalker.vm.HomeScreenViewModelFactory

@Composable
fun HomeScreen(
    modifier: Modifier = Modifier,
    editPet: (Pet?) -> Unit,
    requestWalk: (Pet) -> Unit,
    walkDetails: (Walk) -> Unit
) {
    val context = LocalContext.current
    val navControllerBottomBar = rememberNavController()
    val viewModel = viewModel<HomeScreenViewModel>(factory = HomeScreenViewModelFactory(context))
    var selectedTab by remember { mutableStateOf("pets") }
    Scaffold(
        modifier = modifier,
        bottomBar = {
            Row(modifier = Modifier.fillMaxWidth()) {
                Box(
                    modifier = Modifier
                        .weight(1f)
                        .height(50.dp)
                        .clickable {
                            navControllerBottomBar.navigate("pets")
                            selectedTab = "pets"
                        },
                    contentAlignment = Alignment.Center
                ) {
                    val icon = if (selectedTab == "pets") Icons.Filled.Home else Icons.Outlined.Home
                    val tint =
                        if (selectedTab == "pets") MaterialTheme.colorScheme.primary else MaterialTheme.colorScheme.secondary
                    Icon(icon, "", tint = tint)
                }
                Box(
                    modifier = Modifier
                        .weight(1f)
                        .height(50.dp)
                        .align(Alignment.CenterVertically)
                        .clickable {
                            navControllerBottomBar.navigate("walks")
                            selectedTab = "walks"
                        },
                    contentAlignment = Alignment.Center
                ) {
                    val icon =
                        if (selectedTab == "walks") Icons.Filled.List else Icons.Outlined.List
                    val tint =
                        if (selectedTab == "walks") MaterialTheme.colorScheme.primary else MaterialTheme.colorScheme.secondary
                    Icon(icon, "", tint = tint)
                }
            }
        }
    ) { paddings ->
        NavHost(
            navController = navControllerBottomBar,
            startDestination = "pets",
            route = "bottom_bar"
        ) {
            composable("pets") {
                PetsScreen(
                    modifier = Modifier.padding(paddings),
                    onEdit = editPet,
                    onRequestWalk = requestWalk,
                    viewModel = viewModel
                )
            }
            composable("walks") {
                WalksScreen(
                    modifier = Modifier
                        .fillMaxSize()
                        .padding(paddings),
                    viewModel = viewModel,
                    onSelect = {
                        walkDetails(it)
                    }
                )
            }
        }
    }
}