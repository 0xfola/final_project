use anchor_lang::prelude::*;

#[program]
pub mod attendance {
    use super::*;

    pub fn store_student_data(ctx: Context<StoreStudentData>, ipfs_hash: String, nonce: Vec<u8>, tag: Vec<u8>) -> ProgramResult {
        let student = &mut ctx.accounts.student;
        student.ipfs_hash = ipfs_hash;
        student.nonce = nonce;
        student.tag = tag;
        Ok(())
    }
}

#[account]
pub struct Student {
    pub ipfs_hash: String,
    pub nonce: Vec<u8>,
    pub tag: Vec<u8>,
}

#[derive(Accounts)]
pub struct StoreStudentData<'info> {
    #[account(init, payer = user, space = 128)]
    pub student: Account<'info, Student>,
    #[account(mut)]
    pub user: Signer<'info>,
    pub system_program: Program<'info, System>,
}