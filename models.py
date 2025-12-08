from pydantic import BaseModel, Field
from typing import Dict, Any, Optional

class Proposal(BaseModel):
    """
    Represents a change request from an agent.
    """
    intent_id: str = Field(..., description="Unique identifier for the intent/change")
    delta_patch: str = Field(..., description="The delta patch string")
    provisional: bool = Field(False, description="Whether the change is provisional")

class StateSnapshot(BaseModel):
    """
    Represents a snapshot of the system state.
    """
    data: Dict[str, Any] = Field(default_factory=dict, description="The state data")
    version: Optional[str] = Field(None, description="The version of the state")
