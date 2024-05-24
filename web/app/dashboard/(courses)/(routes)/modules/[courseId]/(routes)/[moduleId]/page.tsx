import React from 'react'
import EditModuleForm from './_components/edit-module-form'
import { auth } from '@/auth'
import getModuleById from '../../_actions/get-module-byid-action'
import { redirect } from 'next/navigation'
import getCourseById from '@/app/dashboard/(courses)/_actions/get-course-byid-action'
import SubmoduleForm
  from "@/app/dashboard/(courses)/(routes)/modules/[courseId]/(routes)/[moduleId]/_components/submodule-form";

const DetailModulePage = async ({
  params
}: {
  params: {
    courseId: string
    moduleId: string
  }
}) => {
  const session = await auth()

  if (!session) {
    return redirect(`/auth/signin?callback=/dashboard/modules/${params.courseId}/${params.moduleId}`)
  }

  const [course, mod] = await Promise.all([
    await getCourseById(session, params.courseId),
    await getModuleById(session, params.courseId, params.moduleId)
  ])

  return (
    <div>
      Detail Module
      <EditModuleForm session={session} course={course} mod={mod}  />

      <div className="mt-5">
        <SubmoduleForm session={session} course={course} mod={mod} />
      </div>
    </div>
  )
}

export default DetailModulePage
