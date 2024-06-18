package com.uvarenko.petwalker.screens

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.sp
import androidx.lifecycle.viewmodel.compose.viewModel
import com.uvarenko.petwalker.vm.StartViewModel
import com.uvarenko.petwalker.vm.StartViewModelFactory

@Composable
fun StartScreen(
    modifier: Modifier,
    onNavigate: (String) -> Unit
) {
    val context = LocalContext.current
    val viewModel = viewModel<StartViewModel>(factory = StartViewModelFactory(context))
    val nav by viewModel.initNav.collectAsState(initial = null)
    if (nav != null) { onNavigate(nav!!) }

    Surface(modifier = modifier) {
        Box(modifier = Modifier.fillMaxSize()) {
            Text(
                modifier = Modifier.align(Alignment.Center),
                fontWeight = FontWeight.ExtraBold,
                fontSize = 34.sp,
                letterSpacing = 10.sp,
                text = "PET\nWALKER",
            )
        }
    }
}