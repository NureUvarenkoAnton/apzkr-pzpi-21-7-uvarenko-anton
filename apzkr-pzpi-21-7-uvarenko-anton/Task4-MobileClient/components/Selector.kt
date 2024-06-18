package com.uvarenko.petwalker.components

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.interaction.MutableInteractionSource
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.lazy.LazyRow
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import com.uvarenko.petwalker.data.UserType

@Composable
fun LabelUi(
    text: String = "",
    selected: Boolean = false,
    // 1. Layout Customization Parameters
    labelTextStyle: TextStyle = TextStyle(
        fontWeight = FontWeight.SemiBold,
        fontSize = 16.sp
    ),
    backgroundColor: Color = Color.White,
    selectedBackgroundColor: Color = Color.Black,
    textColor: Color = Color.Black,
    selectedTextColor: Color = Color.White,
    roundedCornerShapeSize: Dp = 8.dp,
    horizontalPadding: Dp = 16.dp,
    verticalPadding: Dp = 8.dp,
    onClick: () -> Unit = {},
) {
    // 2. Interaction Feedback Removal
    val interactionSource = remember { MutableInteractionSource() }
    Box(
        modifier = Modifier
            .clickable(
                interactionSource = interactionSource,
                indication = null,
                onClick = onClick
            )
            .background(
                // 3. Label Color Depending the selected Value
                if (selected) selectedBackgroundColor else backgroundColor,
                RoundedCornerShape(roundedCornerShapeSize)
            )
            .padding(horizontal = horizontalPadding, vertical = verticalPadding),
        contentAlignment = Alignment.Center
    ) {
        Text(
            text = text,
            // 4. Text Color Depending the selected Value
            color = if (selected) selectedTextColor else textColor,
            style = labelTextStyle
        )
    }
}

@Composable
fun LabelSelectorBar(
    labelItems: List<UserType> = listOf(),
    // 1. LabelSelectorBar Customization Parameters
    barHeight: Dp =  56.dp,
    horizontalPadding: Dp = 8.dp,
    distanceBetweenItems: Dp = 0.dp,
    // 2. LabelUi Customization Parameters
    labelTextStyle: TextStyle = TextStyle(fontWeight = FontWeight.SemiBold, fontSize = 16.sp),
    backgroundColor: Color = Color.White,
    selectedBackgroundColor: Color = Color.Black,
    textColor: Color = Color.Black,
    selectedTextColor: Color = Color.White,
    roundedCornerShapeSize: Dp = 8.dp,
    labelHorizontalPadding: Dp = 16.dp,
    labelVerticalPadding: Dp = 8.dp,
    selected : (UserType) -> Unit
) {
    // 3. Stateful Selection Management:
    val selectedLabel = rememberSaveable { mutableStateOf(labelItems.firstOrNull() ?: UserType.WALKER) }
    LazyRow(
        verticalAlignment = Alignment.CenterVertically,
        modifier = Modifier.height(barHeight)
    ) {
        item { Spacer(modifier = Modifier.width(horizontalPadding)) }
        // 4. Interactive Label Identification
        items(labelItems.size) { label ->
            LabelUi(
                text = labelItems[label].value,
                selected = labelItems[label] == selectedLabel.value,
                labelTextStyle = labelTextStyle,
                backgroundColor = backgroundColor,
                selectedBackgroundColor = selectedBackgroundColor,
                textColor = textColor,
                selectedTextColor = selectedTextColor,
                roundedCornerShapeSize = roundedCornerShapeSize,
                horizontalPadding = labelHorizontalPadding,
                verticalPadding = labelVerticalPadding
            ) {
                selectedLabel.value = labelItems[label]
                selected(selectedLabel.value)
            }
            Spacer(modifier = Modifier.width(distanceBetweenItems))
        }
        item { Spacer(modifier = Modifier.width(horizontalPadding)) }
    }
}
