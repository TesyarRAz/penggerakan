"use client"
import React, { Component, FunctionComponent, useEffect, useState } from "react";
import {
    Dialog,
    DialogContent,
    DialogDescription, DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { FormItem, FormLabel } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { SubModuleStructureType } from "@/types/enums";
import { Label } from "@/components/ui/label";
import { DialogProps } from "@radix-ui/react-dialog";

const ValueItem = ({
    value,
    setValue
}: {
    value: string,
    setValue: (value: string) => void
}) => {

    return (
        <Input value={value} onChange={(e) => setValue(e.target.value)} />
    )

}

// create components record from submodule structure type and value is ValueItem
const resource_components: Record<string, FunctionComponent<{
    value: string,
    setValue: (value: string) => void
}>> = {
    [SubModuleStructureType.Label]: ValueItem,
    [SubModuleStructureType.Text]: ValueItem,
    [SubModuleStructureType.Image]: ValueItem,
    [SubModuleStructureType.Video]: ValueItem,
    [SubModuleStructureType.File]: ValueItem,
    [SubModuleStructureType.Link]: ValueItem,
    [SubModuleStructureType.Quiz]: ValueItem,
}

const StructureModal = ({
    structure,
    onSubmit,
    ...props
}: DialogProps & {
    structure?: SubModuleStructure | null,
    onSubmit: (label: string, value: string) => void
}) => {
    const [label, setLabel] = useState(structure?.label ?? "")
    const [newValue, setNewValue] = useState(structure?.value ?? "")

    useEffect(() => {
        setLabel(structure?.label ?? "")
        setNewValue(structure?.value ?? "")
    }, [structure])

    const handleSubmit = () => {
        onSubmit(label, newValue)

        setLabel("")
        setNewValue("")
    }

    return (
        <Dialog {...props}>
            <DialogContent>
                <DialogHeader>
                    <DialogTitle>Atur Structure</DialogTitle>
                    <DialogDescription>
                        This action cannot be undone. This will permanently delete your account
                        and remove your data from our servers.
                    </DialogDescription>
                </DialogHeader>
                <div>
                    <Label>Title</Label>
                    <Input value={label} onChange={(e) => setLabel(e.target.value)} />
                    <Label>Value</Label>
                    {React.createElement(resource_components[structure?.resource_type ?? SubModuleStructureType.Label], {
                        value: newValue,
                        setValue: setNewValue
                    })}
                </div>
                <DialogFooter className="sm:justify-start">
                    <Button onClick={handleSubmit} type="button" variant="default">
                        Simpan
                    </Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    )
}

export default StructureModal