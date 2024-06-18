package com.uvarenko.petwalker.screens

import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.wrapContentHeight
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.material3.Card
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.uvarenko.petwalker.data.Walk
import com.uvarenko.petwalker.vm.HomeScreenViewModel

@Composable
fun WalksScreen(
    modifier: Modifier = Modifier,
    viewModel: HomeScreenViewModel,
    onSelect: (Walk) -> Unit
) {
    val walks by viewModel.walks.collectAsState()
    Column(modifier = modifier) {
        LazyColumn(modifier = Modifier.fillMaxSize()) {
            items(walks.size) {
                WalkItem(
                    modifier = Modifier
                        .fillMaxWidth()
                        .wrapContentHeight()
                        .clickable { onSelect(walks[it]) }
                        .padding(horizontal = 10.dp),
                    walks[it]
                )
                Spacer(modifier = Modifier.height(10.dp))
            }
        }
    }
}

@Composable
fun WalkItem(modifier: Modifier = Modifier, walk: Walk) {
    Card(
        modifier = modifier
    ) {
        Column(
            modifier = Modifier.padding(10.dp)
        ) {
            Text(text = walk.walkerName, style = MaterialTheme.typography.headlineLarge)
            Text(text = walk.petName, style = MaterialTheme.typography.headlineMedium)
            walk.walkState?.let {
                Text(text = it, style = MaterialTheme.typography.bodySmall)
            }
        }
    }
}