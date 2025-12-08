import requests
from typing import Dict, Any, Optional
from models import Proposal, StateSnapshot

class BSAClient:
    """
    Client for interacting with the BSA Core Service.
    """
    def __init__(self, base_url: str = "http://localhost:8080"):
        self.base_url = base_url.rstrip("/")

    def get_state(self, version: Optional[str] = None) -> Dict[str, Any]:
        """
        Retrieves the canonical state for a given version.
        """
        params = {}
        if version:
            params["version"] = version
        
        response = requests.get(f"{self.base_url}/api/v1/state", params=params)
        response.raise_for_status()
        return response.json()

    def propose_change(self, proposal: Proposal) -> Dict[str, Any]:
        """
        Submits a proposal for a state change.
        """
        response = requests.post(
            f"{self.base_url}/api/v1/propose",
            json=proposal.model_dump()
        )
        response.raise_for_status()
        return response.json()

    def trigger_reconciliation(self):
        """
        Triggers the reconciliation process.
        Note: In the current BSA Core implementation, reconciliation runs automatically
        on a background loop. This method is a placeholder or could be used if
        an explicit trigger endpoint is added in the future.
        """
        print("Triggering reconciliation... (Note: BSA Core runs reconciliation automatically every 5 seconds)")
        # If an endpoint existed:
        # requests.post(f"{self.base_url}/api/v1/reconcile")
