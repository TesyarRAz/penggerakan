import { redirect, useRouter } from 'next/navigation'
import React from 'react'

const MainPage = () => {
  return redirect('/dashboard')
}

export default MainPage
