package com.uvarenko.petwalker

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.ui.Modifier
import androidx.core.os.bundleOf
import androidx.core.view.ViewCompat
import androidx.core.view.WindowCompat
import androidx.core.view.WindowInsetsCompat
import androidx.core.view.WindowInsetsControllerCompat
import androidx.navigation.NavOptions
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.uvarenko.petwalker.screens.HomeScreen
import com.uvarenko.petwalker.screens.PetEditScreen
import com.uvarenko.petwalker.screens.RequestWalkScreen
import com.uvarenko.petwalker.screens.SignInSingUp
import com.uvarenko.petwalker.screens.StartScreen
import com.uvarenko.petwalker.screens.WalkMap
import com.uvarenko.petwalker.ui.theme.PetWalkerTheme

class MainActivity : ComponentActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        makeFullScreen()
        setContent {
            PetWalkerTheme {
                val navController = rememberNavController()
                NavHost(navController = navController, startDestination = "start") {
                    composable("start") {
                        StartScreen(Modifier.fillMaxSize()) {
                            navController.navigate(it)
                        }
                    }
                    composable("signin") {
                        SignInSingUp(Modifier.fillMaxSize(), false) {
                            navController.navigate("home", NavOptions.Builder().setPopUpTo("start", true).build())
                        }
                    }
                    composable("pet_edit") {
                        PetEditScreen(
                            Modifier.fillMaxSize(),
                            it.arguments?.getParcelable("value")
                        ) {
                            navController.popBackStack()
                        }
                    }
                    composable("request_walk") {
                        RequestWalkScreen(
                            Modifier.fillMaxSize(),
                            it.arguments?.getParcelable("value")!!
                        )
                    }
                    composable("walk_details") {
                        WalkMap(
                            Modifier.fillMaxSize(),
                            it.arguments?.getParcelable("value")!!
                        )
                    }
                    composable("home") {
                        HomeScreen(
                            Modifier.fillMaxSize(),
                            editPet = {
                                navController.navigate(
                                    navController.graph.findNode("pet_edit")!!.id,
                                    bundleOf("value" to it)
                                )
                            },
                            requestWalk = {
                                navController.navigate(
                                    navController.graph.findNode("request_walk")!!.id,
                                    bundleOf("value" to it)
                                )
                            },
                            walkDetails = {
                                navController.navigate(
                                    navController.graph.findNode("walk_details")!!.id,
                                    bundleOf("value" to it)
                                )
                            }
                        )
                    }
                }
            }
        }
    }

    private fun makeFullScreen() {
        val windowInsetsController =
            WindowCompat.getInsetsController(window, window.decorView)
        // Configure the behavior of the hidden system bars.
        windowInsetsController.systemBarsBehavior =
            WindowInsetsControllerCompat.BEHAVIOR_SHOW_TRANSIENT_BARS_BY_SWIPE

        // Add a listener to update the behavior of the toggle fullscreen button when
        // the system bars are hidden or revealed.
        ViewCompat.setOnApplyWindowInsetsListener(window.decorView) { view, windowInsets ->
            // You can hide the caption bar even when the other system bars are visible.
            // To account for this, explicitly check the visibility of navigationBars()
            // and statusBars() rather than checking the visibility of systemBars().
            if (windowInsets.isVisible(WindowInsetsCompat.Type.navigationBars())
                || windowInsets.isVisible(WindowInsetsCompat.Type.statusBars())
            ) {
                windowInsetsController.hide(WindowInsetsCompat.Type.systemBars())
            }
            ViewCompat.onApplyWindowInsets(view, windowInsets)
        }
    }

}
