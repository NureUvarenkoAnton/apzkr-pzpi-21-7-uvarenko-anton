package com.uvarenko.petwalker.screens

import androidx.compose.foundation.clickable
import androidx.compose.foundation.interaction.MutableInteractionSource
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.text.selection.LocalTextSelectionColors
import androidx.compose.material3.Button
import androidx.compose.material3.LocalContentColor
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.input.PasswordVisualTransformation
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.lifecycle.viewmodel.compose.viewModel
import com.uvarenko.petwalker.components.LabelSelectorBar
import com.uvarenko.petwalker.data.UserType
import com.uvarenko.petwalker.vm.SignInSignUpViewModel
import com.uvarenko.petwalker.vm.SignInSignUpViewModelFactory
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch

@Preview
@Composable
fun SignInSingUp(
    modifier: Modifier = Modifier,
    isSignUp: Boolean = false,
    onNavigate: () -> Unit = {}
) {
    var isSignUp by remember { mutableStateOf(isSignUp) }
    Scaffold(modifier = modifier) {
        Column(
            modifier = Modifier
                .fillMaxSize()
                .padding(it)
                .padding(20.dp)
        ) {
            Box(
                modifier = Modifier
                    .fillMaxWidth()
                    .weight(1f)
            ) {
                Text(
                    modifier = Modifier,
                    fontWeight = FontWeight.ExtraBold,
                    fontSize = 34.sp,
                    letterSpacing = 10.sp,
                    text = "PET\nWALKER",
                )
            }
            Column(
                modifier = Modifier
                    .fillMaxWidth()
                    .weight(2f)
            ) {
                val context = LocalContext.current
                val viewModel =
                    viewModel<SignInSignUpViewModel>(factory = SignInSignUpViewModelFactory(context))
                var name by remember { mutableStateOf("") }
                var login by remember { mutableStateOf("") }
                var password by remember { mutableStateOf("") }
                var type by remember { mutableStateOf(UserType.WALKER) }
                if (isSignUp)
                    TextField(
                        modifier = Modifier.fillMaxWidth(),
                        value = name,
                        label = { Text("Name") },
                        placeholder = { Text("Name") },
                        onValueChange = { name = it })
                TextField(
                    modifier = Modifier.fillMaxWidth(),
                    value = login,
                    label = { Text("Login") },
                    placeholder = { Text("Login") },
                    onValueChange = { login = it })
                TextField(
                    modifier = Modifier
                        .fillMaxWidth(),
                    value = password,
                    label = { Text("Password") },
                    placeholder = { Text("Password") },
                    visualTransformation = PasswordVisualTransformation(),
                    onValueChange = { password = it }
                )
                // todo fix
                val cs = rememberCoroutineScope()
                if (isSignUp)
                    LabelSelectorBar(
                        labelItems = UserType.entries,
                        backgroundColor = MaterialTheme.colorScheme.background,
                        selectedBackgroundColor = MaterialTheme.colorScheme.primary
                    ) { type = it }
                Button(
                    modifier = Modifier
                        .padding(top = 40.dp)
                        .fillMaxWidth(),
                    onClick = {
                        if (isSignUp) viewModel.signUp(login, password, name, type)
                        else viewModel.signIn(login, password)
                        // todo fix delay
                        cs.launch {
                            delay(500)
                            onNavigate()
                        }
                    }
                ) {
                    Text(text = if (isSignUp) "SignUp" else "LogIn")
                }
                Text(
                    modifier = Modifier
                        .padding(top = 30.dp)
                        .fillMaxWidth()
                        .clickable(
                            interactionSource = remember { MutableInteractionSource() },
                            indication = null
                        ) { isSignUp = !isSignUp },
                    color = MaterialTheme.colorScheme.primary,
                    fontSize = 12.sp,
                    textAlign = TextAlign.Center,
                    text = if (isSignUp) "Already have account?" else "Do not have an account?"
                )
            }
        }
    }
}