package com.uvarenko.petwalker.screens

import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.lifecycle.viewmodel.compose.viewModel
import com.google.android.gms.maps.model.CameraPosition
import com.google.android.gms.maps.model.LatLng
import com.google.maps.android.compose.GoogleMap
import com.google.maps.android.compose.Marker
import com.google.maps.android.compose.MarkerState
import com.google.maps.android.compose.rememberCameraPositionState
import com.uvarenko.petwalker.data.Walk
import com.uvarenko.petwalker.vm.HomeScreenViewModel
import com.uvarenko.petwalker.vm.HomeScreenViewModelFactory

@Composable
fun WalkMap(modifier: Modifier = Modifier, walk: Walk) {
    val context = LocalContext.current
    val viewModel = viewModel<HomeScreenViewModel>(factory = HomeScreenViewModelFactory(context))
    val kyiv = LatLng(50.450001, 30.523333)
    val cameraPositionState = rememberCameraPositionState {
        position = CameraPosition.fromLatLngZoom(kyiv, 10f)
    }
    GoogleMap(
        modifier = modifier,
        cameraPositionState = cameraPositionState
    ) {
        val coordinates by viewModel.walkPosition.collectAsState()
        if (coordinates != null) {
            Marker(
                state = MarkerState(position = LatLng(coordinates!!.first.toDouble(), coordinates!!.second.toDouble())),
                title = "Pet"
            )
        }
    }
}