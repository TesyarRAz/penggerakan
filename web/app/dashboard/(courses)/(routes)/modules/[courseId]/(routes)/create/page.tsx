import getCourseById from '@/app/dashboard/(courses)/_actions/get-course-byid-action'
import { auth } from '@/auth'
import { notFound, redirect } from 'next/navigation'
import React from 'react'
import CreateModuleForm from './_components/create-module-form'

const CreateModulePage = async ({
    params
  }: {
    params: {
      courseId: string
    }
  }) => {
    const session = await auth()
  
    if (!session) {
      return redirect(`/auth/signin?callback=/dashboard/modules/${params.courseId}/create`)
    }
  
    const course = await getCourseById(session, params.courseId)
    if (!course) {
      return notFound()
    }
  
  return (
    <CreateModuleForm session={session} course={course} />
  )
}

export default CreateModulePage
