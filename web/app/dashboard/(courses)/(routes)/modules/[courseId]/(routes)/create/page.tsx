import { authOptions } from '@/app/api/auth/[...nextauth]/route'
import getCourseById from '@/app/dashboard/(courses)/_actions/get-course-byid-action'
import { getServerSession } from 'next-auth'
import React from 'react'
import CreateModuleForm from './_components/create-module-form'
import { redirect } from 'next/navigation'

const CreateModulePage = async ({
  params
}: {
  params: {
    courseId: string
  }
}) => {
  const session = await getServerSession(authOptions)

  if (!session) {
    return redirect('/auth/signin')
  }

  const course = await getCourseById(session, params.courseId)

  return (
    <CreateModuleForm course={course} />
  )
}

export default CreateModulePage
