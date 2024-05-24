import React from 'react'
import LoginForm from './_components/login-form'
import { auth } from '@/auth'
import { redirect } from 'next/navigation'

const LoginPage = async () => {
  const session = await auth()

  if (session && !session?.token?.error) {
    return redirect("/dashboard")
  }

  return (
    <LoginForm/>
  )
}

export default LoginPage
