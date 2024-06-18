package com.uvarenko.petwalker.data

import android.content.Context

class LocalStore(context: Context) {

    private val preferences = context.getSharedPreferences("store", Context.MODE_PRIVATE)

    fun saveToken(token: String) {
        preferences.edit().putString("token", token).apply()
    }

    fun getToken() = preferences.getString("token", null)

}