import React from 'react'
import EditModuleForm from './_components/edit-module-form'
import { auth } from '@/auth'
import getModuleById from '../../_actions/get-module-byid-action'
import { redirect } from 'next/navigation'
import getCourseById from '@/app/dashboard/(courses)/_actions/get-course-byid-action'
import SubModuleForm
  from "@/app/dashboard/(courses)/(routes)/modules/[courseId]/(routes)/[moduleId]/_components/submodule-form";
import getSubModules from '../../_actions/get-submodules-action'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import BreadcumbLayout from '@/components/layouts/breadcumb-layout'

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

  const [course, mod, submodules] = await Promise.all([
    await getCourseById(session, params.courseId),
    await getModuleById(session, params.courseId, params.moduleId),
    await getSubModules(session, params.moduleId),
  ])

  return (
    <BreadcumbLayout
      title="Detail Module"
    >
      <EditModuleForm session={session} course={course} mod={mod} />

      <div className="mt-5">
        <h2>SubModules</h2>
        <Tabs>
          <TabsList className="flex space-x-3">
            {submodules.data.map((submodule) => (
              <TabsTrigger key={submodule.id} value={submodule.id}>{submodule.name}</TabsTrigger>
            ))}
            <TabsTrigger value="new">New Sub Module</TabsTrigger>
          </TabsList>
          {submodules.data.map((submodule) => (
            <TabsContent key={submodule.id} value={submodule.id}>
              <SubModuleForm session={session} submodule={submodule} mod={mod} />
            </TabsContent>
          ))}
          <TabsContent value="new">
            <SubModuleForm session={session} mod={mod} />
          </TabsContent>
        </Tabs>
      </div>
    </BreadcumbLayout>
  )
}

export default DetailModulePage
