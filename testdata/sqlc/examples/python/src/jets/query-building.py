# Code generated by sqlc. DO NOT EDIT.
from typing import AsyncIterator, Optional

import sqlalchemy
import sqlalchemy.ext.asyncio

from jets import models


COUNT_PILOTS = """-- name: count_pilots \\:one
SELECT COUNT(*) FROM pilots
"""


DELETE_PILOT = """-- name: delete_pilot \\:exec
DELETE FROM pilots WHERE id = :p1
"""


LIST_PILOTS = """-- name: list_pilots \\:many
SELECT id, name FROM pilots LIMIT 5
"""


class AsyncQuerier:
    def __init__(self, conn: sqlalchemy.ext.asyncio.AsyncConnection):
        self._conn = conn

    async def count_pilots(self) -> Optional[int]:
        row = (await self._conn.execute(sqlalchemy.text(COUNT_PILOTS))).first()
        if row is None:
            return None
        return row[0]

    async def delete_pilot(self, *, id: int) -> None:
        await self._conn.execute(sqlalchemy.text(DELETE_PILOT), {"p1": id})

    async def list_pilots(self) -> AsyncIterator[models.Pilot]:
        result = await self._conn.stream(sqlalchemy.text(LIST_PILOTS))
        async for row in result:
            yield models.Pilot(
                id=row[0],
                name=row[1],
            )