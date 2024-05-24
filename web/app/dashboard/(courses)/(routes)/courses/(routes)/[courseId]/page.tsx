import React from 'react'
import getCourseById from '../../../../_actions/get-course-byid-action'
import { redirect } from 'next/navigation'
import { auth } from '@/auth'

const EditCoursePage = async ({
  params
}: {
  params: {
    courseId: string
  }
}) => {
  const session = await auth()

  if (!session) {
    return redirect(`/auth/signin?callback=/dashboard/courses/${params.courseId}`)
  }

  const course = await getCourseById(session, params.courseId)

  return (
    <div>
      <h2>Detail Page</h2>
      
    </div>
  )
}

export default EditCoursePage
