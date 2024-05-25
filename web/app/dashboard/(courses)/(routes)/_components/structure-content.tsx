import { Button } from '@/components/ui/button'
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from '@/components/ui/collapsible'
import { cn } from '@/lib/utils'
import Link from 'next/link'
import React, { useState } from 'react'
import { FaSortDown, FaSortUp } from 'react-icons/fa'

const ListItem = ({
    structure,
    ...props
}: React.HTMLAttributes<HTMLDivElement> & {
    structure: SubModuleStructure,
}) => {
    const [open, setOpen] = useState(true)

    const url = (structure?.value !== null && structure?.value !== "") ? structure?.value ?? "#" : "#"

    return (
        <Collapsible open={open} onOpenChange={setOpen}>
            <div className="ml-2" {...props}>
                <div className="flex justify-between">
                    <Link href={url}>{structure.label}</Link>
                    <div>
                        <CollapsibleTrigger asChild>
                            <Button variant="ghost">
                                <FaSortUp className={cn("h-4 w-4", open ? "hidden" : "")} />
                                <FaSortDown className={cn("h-4 w-4", !open ? "hidden" : "")} />
                            </Button>
                        </CollapsibleTrigger>
                    </div>
                </div>

                <div className="mb-3">
                    <CollapsibleContent>
                        {structure.children?.map((item, index) => (
                            <ListItem
                                key={index}
                                structure={item} />
                        ))}
                    </CollapsibleContent>
                </div>
            </div>
        </Collapsible>
    )
}

const StructureContent = ({
    structures,
    ...props
}: React.HTMLAttributes<HTMLDivElement> & {
    structures: SubModuleStructure[],
}) => {
    return (
        <div>
            {structures.map((structure, index) => (
                <ListItem
                    key={index}
                    structure={structure} />
            ))}
        </div>
    )
}

export default StructureContent
