"use client"
import React, {useState} from 'react'
import {Session} from "next-auth";
import {Button} from "@/components/ui/button";

interface StructureItem {

}

const SubModuleForm = ({
   session,
   course,
   mod,
}: {
    session: Session,
    course: CourseResponse,
    mod: ModuleResponse,
}) => {
    const [items, setItems] = useState<StructureItem[]>([])

    return (
        <div>
            {items.map((item, index) => (
                <div
                key={index}
                >

                </div>
            ))}
            <Button onClick={() => setItems([...items, {}])}>Add</Button>
        </div>
    )
}

export default SubModuleForm
