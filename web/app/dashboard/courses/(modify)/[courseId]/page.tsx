import { authOptions } from '@/app/api/auth/[...nextauth]/route'
import { getServerSession } from 'next-auth'
import React from 'react'

const EditCoursePage = async ({
  params
}: {
  params: {
    courseId: string
  }
}) => {
  const session = await getServerSession(authOptions)
  const course = await getCourseById(session, params.courseId)

  return (
    <div>
      <h2>Detail Page</h2>
      
    </div>
  )
}

export default EditCoursePage
