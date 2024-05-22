import { authOptions } from '@/app/api/auth/[...nextauth]/route'
import { getServerSession } from 'next-auth'
import React from 'react'
import getCourseById from '../../../_actions/get-course-byid-action'
import { redirect } from 'next/navigation'

const EditCoursePage = async ({
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
    <div>
      <h2>Detail Page</h2>
      
    </div>
  )
}

export default EditCoursePage
